package healthPinger

func UpdateTaskInfo_burstTime(burstTime float64, provider string, model string, version string) {
	prevTaskInfo := model_info[provider+"@"+model][version]
	var nextTaskInfo TaskInfo

	if prevTaskInfo.AverageInferenceTime == 0 {
		nextTaskInfo = TaskInfo{
			AverageInferenceTime: float32(burstTime),
			LoadedAmount:         prevTaskInfo.LoadedAmount,
		}
	} else {
		nextTaskInfo = TaskInfo{
			AverageInferenceTime: (float32(burstTime) + prevTaskInfo.AverageInferenceTime) / 2,
			LoadedAmount:         prevTaskInfo.LoadedAmount,
		}
	}

	model_info[provider+"@"+model][version] = nextTaskInfo
}

func UpdateTaskInfo_start(provider string, model string, version string) {
	prevTaskInfo := model_info[provider+"@"+model][version]

	model_info[provider+"@"+model][version] = TaskInfo{
		AverageInferenceTime: prevTaskInfo.AverageInferenceTime,
		LoadedAmount:         prevTaskInfo.LoadedAmount + 1,
	}
}
func UpdateTaskInfo_end(provider string, model string, version string) {
	prevTaskInfo := model_info[provider+"@"+model][version]

	model_info[provider+"@"+model][version] = TaskInfo{
		AverageInferenceTime: prevTaskInfo.AverageInferenceTime,
		LoadedAmount:         prevTaskInfo.LoadedAmount - 1,
	}
}

func UpdateModel(provider string, model string, version string) {
	model_info[provider+"@"+model][version] = TaskInfo{AverageInferenceTime: 0, LoadedAmount: 0}
}
