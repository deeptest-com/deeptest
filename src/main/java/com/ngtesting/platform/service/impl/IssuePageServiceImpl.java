package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageDao;
import com.ngtesting.platform.service.IssuePageService;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class IssuePageServiceImpl extends BaseServiceImpl implements IssuePageService {
    @Autowired
    UserService userService;

	@Autowired
    IssuePageDao pageDao;



}
