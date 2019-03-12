package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.UserService;
import com.ngtesting.platform.servlet.PrivOrg;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping(value = Constant.API_PATH_CLIENT + "/user")
public class UserAction extends BaseAction {
    @Autowired
    private UserService userService;
    @Autowired
    private UserDao userDao;

    @PostMapping(value = "getUsers")
    @PrivPrj
    public Map<String, Object> getUsers(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

        List <TstUser> vos = userService.getProjectUsers(orgId, prjId);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "search")
    @PrivOrg(perms = {"org_org:*"})
    public Map<String, Object> search(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        String keywords = json.getString("keywords");
        List<Integer> exceptIds = json.getObject("exceptIds", List.class);

        List users = userService.searchPrjUser(user.getDefaultPrjId(), keywords, exceptIds);

        List<Object> vos = new ArrayList<>();
        vos.addAll(users);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
