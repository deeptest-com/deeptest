package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.AuthDao;
import com.ngtesting.platform.dao.ProjectConfigDao;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.service.intf.AuthService;
import com.ngtesting.platform.service.intf.ProjectConfigService;
import com.ngtesting.platform.service.intf.PushSettingsService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ProjectConfigServiceImpl extends BaseServiceImpl implements ProjectConfigService {

	private static final Log log = LogFactory.getLog(ProjectConfigServiceImpl.class);

	private ProjectDao projectDao;
    @Autowired
    private ProjectConfigDao projectConfigDao;

    @Autowired
    AuthService authService;
    @Autowired
    AuthDao authDao;

    @Autowired
    private PushSettingsService pushSettingsService;

	@Override
	public TstProject get(Integer id) {
		if (id == null) {
			return null;
		}
		TstProject po = projectDao.get(id);

		return po;
	}
}
