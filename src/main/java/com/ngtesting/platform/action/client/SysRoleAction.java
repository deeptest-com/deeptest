package com.ngtesting.platform.action.client;

import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.intf.SysRoleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "sys_role/")
public class SysRoleAction extends BaseAction {
	@Autowired
    SysRoleService roleService;


}
