package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.UserService;
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

    @PostMapping(value = "getProjectUsers") // 用例筛选页面，通过用户筛选
    @PrivPrj(perms = {"belongs_to:project"})
    public Map<String, Object> getProjectUsers(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        List <TstUser> vos = userService.getProjectUsers(prjId);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "searchProjectUser") // 任务指定用户时调用
    @PrivPrj(perms = {"test_plan:maintain"})
    public Map<String, Object> searchProjectUser(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        String keywords = json.getString("keywords");
        List<Integer> exceptIds = json.getObject("exceptIds", List.class);

        List users = userService.searchProjectUser(user.getDefaultPrjId(), keywords, exceptIds);

        List<Object> vos = new ArrayList<>();
        vos.addAll(users);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
