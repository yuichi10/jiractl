package usecase

// ActiveSprintIssues get apis
func ActiveSprintIssues(api IJiraAPIAccess, presenter ISprintPresenter) {
	api.GetActiveSprintIssues()
}
