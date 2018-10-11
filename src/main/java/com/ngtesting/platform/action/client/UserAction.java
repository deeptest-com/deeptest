package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/user")
public class UserAction extends BaseAction {
    @Autowired
    private UserService userService;
    @Autowired
    private UserDao userDao;

    @PostMapping(value = "listLastest")
    @ResponseBody
    public Map<String, Object> listLastest(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer projectId = user.getDefaultPrjId();

        List<TstUser> users = userDao.getProjectUsers(projectId, 10);

        ret.put("data", users);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "getUsers")
    @ResponseBody
    public Map<String, Object> getUsers(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();

        Integer projectId = json.getInteger("projectId");

        List <TstUser> vos = userService.getProjectUsers(orgId, projectId);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "search")
    @ResponseBody
    public Map<String, Object> search(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        String keywords = json.getString("keywords");
        JSONArray exceptIds = json.getJSONArray("exceptIds");

        String ids = "";
        if (exceptIds != null && exceptIds.size() > 0) {
            int i = 0;
            for (Object item : exceptIds.toArray()) {
                if (i++ > 0) {
                    ids += ",";
                }
                ids += item.toString();
            }
        }

        List users = userService.search(user.getDefaultOrgId(), keywords, ids);

        List<Object> vos = new ArrayList<>();
        vos.addAll(users);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
