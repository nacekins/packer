package vagrant

type StepUp struct{}

func (s *StepUp) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(Driver)

	err := driver.Up()
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *StepUp) Cleanup(state multistep.StateBag) {
}