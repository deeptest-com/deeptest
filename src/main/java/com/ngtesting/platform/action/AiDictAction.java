package com.ngtesting.platform.action;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.AiDictService;
import com.ngtesting.platform.util.AuthPassport;
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
@RequestMapping(Constant.API_PATH_CLIENT + "ai_dict/")
public class AiDictAction extends BaseAction {
	@Autowired
	AiDictService dictService;

	@AuthPassport(validate = false)
	@RequestMapping(value = "importDict", method = RequestMethod.GET)
	@ResponseBody
	public Map<String, Object> importDict(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();

		dictService.dictPers();

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = false)
	@RequestMapping(value = "load", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> load(HttpServletRequest request, @RequestBody String json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Map<String, List<String>> map = dictService.get(json);

		ret.put("data", map);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
