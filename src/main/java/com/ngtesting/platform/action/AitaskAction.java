package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.AiTestTask;
import com.ngtesting.platform.service.*;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.vo.*;
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
@RequestMapping(Constant.API_PATH_CLIENT + "aitask/")
public class AitaskAction extends BaseAction {
	@Autowired
	ProjectService projectService;
	@Autowired
    AiTestTaskService aitaskService;

    @Autowired
    AiAsrLangModelService aiAsrLangModelService;
    @Autowired
    AiAudioTypeService aiAudioTypeService;
    @Autowired
    AiProductBranchService aiProductBranchService;
    @Autowired
    AiTestEnvService aiTestEnvService;
    @Autowired
    AiTestTypeService aiTestTypeService;
    @Autowired
    AiTestSetService aiTestSetService;
    @Autowired
    JenkinsService jenkinsSetService;

    @Autowired
    AiTestScheduleService aiTestScheduleService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        Long orgId = json.getLong("orgId");
		Long projectId = json.getLong("projectId");

		List<AiTestTask> ls = aitaskService.query(projectId);
        List<AiTestTaskVo> vos = aitaskService.genVos(ls);

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
        Long projectId = userVo.getDefaultPrjId();
        Long caseId = json.getLong("id");

        AiTestTaskVo vo = aitaskService.getById(caseId);

        List<AiAsrLangModelVo> asrLangModelVos = aiAsrLangModelService.listAsrLangModelVo(projectId);
        List<AiAudioTypeVo> aiAudioTypeVos = aiAudioTypeService.listAudioTypeVo(projectId);
        List<AiProductBranchVo> aiProductBranchVos = aiProductBranchService.listForProductBranchVo(projectId);
        List<AiTestEnvVo> aiTestEnvVos = aiTestEnvService.listTestEnvVo(projectId);
        List<AiTestTypeVo> aiTestTypeVos = aiTestTypeService.listTestTypeVo(projectId);

        List<AiTestSetVo> aiTestSetVos = aiTestSetService.listTestSetVo(projectId);

        ret.put("asrLangModelVos", asrLangModelVos);
        ret.put("aiAudioTypeVos", aiAudioTypeVos);
        ret.put("aiProductBranchVos", aiProductBranchVos);
        ret.put("aiTestEnvVos", aiTestEnvVos);
        ret.put("aiTestTypeVos", aiTestTypeVos);
        ret.put("aiTestSetVos", aiTestSetVos);

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

		AiTestTask po = aitaskService.save(json, userVo);
		AiTestTaskVo caseVo = aitaskService.genVo(po);

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

		AiTestTask testCasePo = aitaskService.renamePers(json, userVo);
        aitaskService.updateParentIfNeededPers(testCasePo.getpId());
		AiTestTaskVo caseVo = aitaskService.genVo(testCasePo);

		ret.put("data", caseVo);
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

		AiTestTask testCase = aitaskService.delete(id, userVo);
		aitaskService.updateParentIfNeededPers(testCase.getpId());

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "move", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> move(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        Long srcId = json.getLong("srcId");
        Long parentId = aitaskService.getById(srcId).getpId();
        Long targetId = json.getLong("targetId");
        AiTestTaskVo vo = aitaskService.movePers(json, userVo);

        aitaskService.updateParentIfNeededPers(parentId);
        aitaskService.updateParentIfNeededPers(targetId);

		ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "run", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> run(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        Long taskId = json.getLong("taskId");
        AiTestTaskVo vo = aitaskService.getById(taskId);

        String result = jenkinsSetService.execute(vo);

        ret.put("json", result);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
