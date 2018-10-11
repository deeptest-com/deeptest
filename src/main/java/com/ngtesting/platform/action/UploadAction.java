package com.ngtesting.platform.action;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.utils.FileUtil;
import org.apache.commons.io.FilenameUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletRequest;
import java.text.DecimalFormat;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;


@Controller
@RequestMapping(Constant.API_PATH + "/")
public class UploadAction extends BaseAction {
	@Autowired
	UserService userService;

	@ResponseBody
	@PostMapping("/uploadSingle")
	public Map<String, Object> uploadSingle(
            @RequestParam("file") MultipartFile file, HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String origName = file.getOriginalFilename();
		String extName = FilenameUtils.getExtension(origName);
		String fileName = UUID.randomUUID().toString() + "." + extName;

		String uploadPath = FileUtil.SaveFile(file, "data/", fileName);

		ret.put("origName", origName);
		ret.put("uploadPath", uploadPath);

		float flt = Float.parseFloat(String.valueOf(file.getSize()));
		String fileSize = new DecimalFormat("##0.00").format(flt / 1000 / 1000);
		ret.put("fileSize", fileSize + 'M');

		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
