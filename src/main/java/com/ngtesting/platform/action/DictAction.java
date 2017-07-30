package com.ngtesting.platform.action;

import com.ngtesting.platform.service.PlanService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;

@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "dict/")
public class DictAction extends BaseAction {
	@Autowired
	PlanService planService;
	
	@AuthPassport(validate = false)
	@RequestMapping(value = "dict", method = RequestMethod.GET)
	@ResponseBody
	public Map<String, Object> dict(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		planService.dictPers();

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
