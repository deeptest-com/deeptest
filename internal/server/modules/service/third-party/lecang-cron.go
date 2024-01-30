package third_party

type LecangCron struct {
}

func (t *LecangCron) Run(options map[string]interface{}) (f func() error) {
	f = func() error {
		return nil
	}
	return
}

func (t *LecangCron) CallBack(options map[string]interface{}, err error) func() {
	return func() {

	}
}
