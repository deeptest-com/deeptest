package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.Page;
import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

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

    @Autowired
    private OrgService orgService;

    @Autowired
    OrgGroupUserRelationService orgGroupUserRelationService;

    @Autowired
    private ProjectService projectService;

    @Autowired
    SysPrivilegeService sysPrivilegeService;
    @Autowired
    OrgPrivilegeService orgRolePrivilegeService;
    @Autowired
    CasePropertyService casePropertyService;
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;
    @Autowired
    PushSettingsService pushSettingsService;

    @PostMapping(value = "list")
    @ResponseBody
    public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();

        String keywords = json.getString("keywords");
        Boolean disabled = json.getBoolean("disabled");
        Integer pageNum = json.getInteger("page");
        Integer pageSize = json.getInteger("pageSize");

        Page page = PageHelper.startPage(pageNum, pageSize);
        List<TstUser> users = userService.list(orgId, keywords, disabled, pageNum, pageSize);

        ret.put("total", page.getTotal());
        ret.put("data", users);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

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

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer userId = json.getInteger("id");

        List<TstOrgGroupUserRelation> relations = orgGroupUserRelationService.listRelationsByUser(orgId, userId);

        if (userId == null) { // 新增用户
            ret.put("user", new TstUser());
            ret.put("relations", relations);
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
            return ret;
        }

        TstUser po = userService.get(userId);
        if (userNotInOrg(user.getId(), po.getDefaultOrgId())) { // 获取的非当前组织用户
            return authFail();
        }

        ret.put("user", po);
        ret.put("relations", relations);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "invite")
    @ResponseBody
    public Map<String, Object> invite(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        TstUser vo = JSON.parseObject(JSON.toJSONString(json.get("user")), TstUser.class);
        List<TstOrgGroupUserRelation> relations = (List<TstOrgGroupUserRelation>) json.get("relations");
        TstUser po = userService.invite(user, vo, relations);

        if (po == null) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "邮箱已加入当前组织");
            return ret;
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    // 管理员修改用户信息
    @PostMapping(value = "update")
    @ResponseBody
    public Map<String, Object> update(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();

        TstUser vo = JSON.parseObject(JSON.toJSONString(json.get("user")), TstUser.class);

        if (userNotInOrg(vo.getId(), user.getDefaultOrgId())) {
            return authFail();
        }

        TstUser existUser = userService.getByEmail(vo.getEmail());
        if (existUser != null && existUser.getId() != vo.getId()) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "邮箱已被占用");
            return ret;
        }

        TstUser po = userService.update(vo);

        List<TstOrgGroupUserRelation> relations = (List<TstOrgGroupUserRelation>) json.get("relations");
        orgGroupUserRelationService.saveRelationsForUser(orgId, vo.getId(), relations);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        pushSettingsService.pushUserSettings(po);
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
