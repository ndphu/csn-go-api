package service

import "time"

type ScheduleService struct {
	accountService *AccountService
}

var scheduleService *ScheduleService

func GetScheduleService() (*ScheduleService, error) {
	if scheduleService == nil {
		as, err := GetAccountService()
		if err != nil {
			return nil, err
		}
		scheduleService = &ScheduleService{
			accountService: as,
		}
	}

	return scheduleService, nil
}

func (s *ScheduleService) Start() {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.accountService.UpdateAccountCache()
				s.accountService.UpdateAllAccountQuota()
			}
		}
	}()
}
