package schedule

import (
	"github.com/go-redis/redis"
	"walrus/models"
  "time"
  "github.com/satori/go.uuid"
  "fmt"
  "errors"
  "walrus/utils"
  "walrus/constants"
)

type redisScheduleQueue struct {
	client *redis.Client
}

func newJob(jobType string, payload string, runAt int64) models.Job {
  id := uuid.NewV4()
  return models.Job { Id: id.String(), Type: jobType, Payload: payload, RunAt: runAt }
}

func (r *redisScheduleQueue) Add(jobType string, payload string, runAfterSeconds time.Duration) (string, error) {
  runAt := time.Now().Add(runAfterSeconds * time.Second).UnixNano()
  job := newJob(jobType, payload, runAt)

  jobJs, err := utils.ToJson(job)
  if err != nil {
    return "", errors.New(fmt.Sprintf("Unable to serialize job, Error: %s", err))
  }

  jobKeyField := utils.GetJobKeyField(job.Id)
  hSetCmd := r.client.HSet(constants.JOBS_MAP, jobKeyField, jobJs)
  if hSetCmd.Err() != nil {
    return "", errors.New(fmt.Sprintf("Unable to add job to queue, Error: %s", hSetCmd.Err()))
  }

  zAddCmd := r.client.ZAdd(constants.SCHEDULER_QUEUE, redis.Z{
    Score: float64(runAt),
    Member: jobJs,
  })

  if zAddCmd.Err() != nil {
    return "", errors.New(fmt.Sprintf("Unable to add job to queue, Error: %s", zAddCmd.Err()))
  }

	return job.Id, nil
}

func (r *redisScheduleQueue) getSavedJob(jobId string) (*models.Job, error) {
  jobKeyField := utils.GetJobKeyField(jobId)
  hGetCmd := r.client.HGet(constants.JOBS_MAP, jobKeyField)
  if hGetCmd.Err() != nil {
    return nil, errors.New(fmt.Sprintf("Could not fetch job, Error: %s",hGetCmd.Err()))
  }
  jobJs := hGetCmd.Val()
  job, err := utils.ToJob(jobJs)
  if err != nil {
    return nil, errors.New(fmt.Sprintf("Unable to Unmarshall job with id: %s, Error: %s", jobId, err))
  }
  return job, nil
}

func (r *redisScheduleQueue) Update(jobId string, payload string) error {

  job, err := r.getSavedJob(jobId)
  if err != nil {
    return errors.New(fmt.Sprintf("Unable to get job : %s", err))
  }
  job.Payload = payload
  jobJs, jsonErr := utils.ToJson(job)
  if jsonErr != nil {
    panic("Unable to serialize back to json in update")
  }

  hSetCmd := r.client.HSet(constants.JOBS_MAP, utils.GetJobKeyField(jobId), jobJs)
  if hSetCmd.Err() != nil {
    return errors.New(fmt.Sprintf("Unable to update job, Error: %s", hSetCmd.Err()))
  }

	return nil
}

func (r *redisScheduleQueue) Delete(jobId string) error {
  job, err := r.getSavedJob(jobId)
  if err != nil {
    return errors.New(fmt.Sprintf("Unable to get job to delete: %s", err))
  }

  jobJs, jsonErr := utils.ToJson(job)
  if jsonErr != nil {
    panic("Unable to serialize back to json in delete")
  }

  hDelCmd := r.client.HDel(constants.JOBS_MAP, utils.GetJobKeyField(jobId), jobJs)
  if hDelCmd.Err() != nil {
    return errors.New(fmt.Sprintf("Unable to delete job, Error: %s", hDelCmd.Err()))
  }

	return nil
}



func newRedisScheduleQueue() Schedule {
	r := redisScheduleQueue{}
  options := utils.LoadRedisOptions()
	r.client = redis.NewClient(options)

	return &r
}
