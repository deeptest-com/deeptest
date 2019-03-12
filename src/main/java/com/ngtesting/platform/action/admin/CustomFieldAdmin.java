package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.CustomField;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CustomFieldService;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping(Constant.API_PATH_ADMIN + "custom_field/")
public class CustomFieldAdmin extends BaseAction {
    @Autowired
    CustomFieldService customFieldService;

    @RequestMapping(value = "list", method = RequestMethod.POST)
    public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer orgId = user.getDefaultOrgId();

        String applyTo = json.getString("applyTo");
        String keywords = json.getString("keywords");

        List<CustomField> vos = customFieldService.list(orgId, applyTo, keywords); // 总是取当前活动org的，不需要再鉴权

        Map inputMap = customFieldService.inputMap();
        Map typeMap = customFieldService.typeMap();

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("data", vos);
        ret.put("inputMap", inputMap);
        ret.put("typeMap", typeMap);
        return ret;
    }

    @RequestMapping(value = "get", method = RequestMethod.POST)
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer orgId = user.getDefaultOrgId();

        Integer id = json.getInteger("id");

        Map<String, Map> inputMap = customFieldService.fetchInputMap();
        Map typeMap = customFieldService.typeMap();

        List<String> formatList = customFieldService.listFormat();
        List<String> applyToList = customFieldService.listApplyTo();

        CustomField vo = null;
        if (id == null) {
            vo = new CustomField();
            vo.setColCode(customFieldService.getLastUnusedColumn(orgId));
        } else {
            vo = customFieldService.getDetail(id, orgId);
        }

        if (vo == null) { // 当对象不是默认org的，此处为空
            return authorFail();
        }

        if (vo.getColCode() == null) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "自定义字段不能超过20个");
        }

        ret.put("data", vo);
        ret.put("inputMap", inputMap);
        ret.put("typeMap", typeMap);

        ret.put("formatList", formatList);
        ret.put("applyToList", applyToList);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "save", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer orgId = user.getDefaultOrgId();

        CustomField customField = JSON.parseObject(JSON.toJSONString(json.get("model")), CustomField.class);

        CustomField po = customFieldService.save(customField, orgId);
        if (po == null) { // 当对象不是默认org的，update的结果会返回空
            return authorFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "delete", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer orgId = user.getDefaultOrgId();

        Integer id = json.getInteger("id");

        Boolean result = customFieldService.delete(id, orgId);
        if (!result) { // 当对象不是默认org的，结果会返回false
            return authorFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "changeOrder", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer orgId = user.getDefaultOrgId();

        Integer id = json.getInteger("id");
        String act = json.getString("act");

        String applyTo = json.getString("applyTo");
        String keywords = json.getString("keywords");

        Boolean result = customFieldService.changeOrderPers(id, act, orgId, applyTo);
        if (!result) { // 当对象不是默认org的，结果会返回false
            return authorFail();
        }

        List<CustomField> vos = customFieldService.list(orgId, applyTo, keywords);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

}
