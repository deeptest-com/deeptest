package com.ngtesting.platform.action.client;

import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.intf.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/project_config")
public class ProjectConfigAction extends BaseAction {

    @Autowired
    private ProjectConfigService projectConfigService;

}
