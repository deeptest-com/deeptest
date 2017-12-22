package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.service.CaseService;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.TestCaseVo;
import com.ngtesting.platform.vo.UserVo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "case/")
public class CaseAction extends BaseAction {
	@Autowired
    CaseService caseService;
	@Autowired
	CustomFieldService customFieldService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long projectId = json.getLong("projectId");

		List<TestCase> ls = caseService.query(projectId);
        List<TestCaseVo> vos = caseService.genVos(ls, false);

        List<CustomFieldVo> customFieldList = customFieldService.listForCaseByProject(projectId);

        ret.put("data", vos);
		ret.put("customFields", customFieldList);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "queryForSelection", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> queryForSelection(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long projectId = json.getLong("projectId");
		Long runId = json.getLong("runId");

        List<TestCaseVo> vos = caseService.queryForSelection(projectId, runId);

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @AuthPassport(validate = true)
    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Long orgId = userVo.getDefaultOrgId();
        Long caseId = json.getLong("id");

        TestCaseVo vo = caseService.getById(caseId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TestCase po = caseService.save(json, userVo.getId());
		TestCaseVo caseVo = caseService.genVo(po, true);

		ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "rename", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> rename(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TestCase po = caseService.renamePers(json, userVo.getId());
		TestCaseVo caseVo = caseService.genVo(po);

		ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "move", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> move(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        TestCaseVo vo = caseService.movePers(json, userVo.getId());

		ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @AuthPassport(validate = true)
    @RequestMapping(value = "delete", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        Long id = json.getLong("id");

        UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        TestCase po = caseService.delete(id, userVo.getId());

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@AuthPassport(validate = true)
	@RequestMapping(value = "saveField", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveField(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TestCase po = caseService.saveField(json);
        TestCaseVo caseVo = caseService.genVo(po);

		ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "changeContentType", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeContentType(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long id = json.getLong("id");
        String contentType = json.getString("contentType");

		TestCase po = caseService.changeContentTypePers(id, contentType);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
