function dp_test(name, cb) {
	try {
		cb();
	} catch(err){
		check(false, name, err)
		return
	}

	check(true, name, '')
}