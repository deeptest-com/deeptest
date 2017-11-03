package com.ngtesting.platform.action;

import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.util.Constant.RespCode;
import com.ngtesting.platform.vo.UserVo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.multipart.MultipartHttpServletRequest;
import org.springframework.web.multipart.commons.CommonsMultipartResolver;

import javax.servlet.http.HttpServletRequest;
import java.io.File;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "/")
public class UploadAction extends BaseAction {
	@Autowired
	UserService userService;
	@Autowired
	RelationOrgGroupUserService orgGroupUserService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "uploadSingle", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> uploadSingle(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

		// 创建一个通用的多部分解析器
		CommonsMultipartResolver multipartResolver = new CommonsMultipartResolver(request.getSession()
				.getServletContext());
		if (multipartResolver.isMultipart(request)) {
			String uploadDist = request.getSession().getServletContext().getRealPath("/")
					+ Constant.FTP_UPLOAD_DIR + File.separator;

			// 图片相关处理
			MultipartHttpServletRequest multiRequest = (MultipartHttpServletRequest) request;
			MultipartFile icon = multiRequest.getFile("eicon");
			MultipartFile pic = multiRequest.getFile("epic");
			String picName = UUID.randomUUID().toString() + ".png";

			try {
				// 图标文件
				if (icon != null) {
					File icon_local = new File(uploadDist + "icon" + File.separator + picName);
					icon.transferTo(icon_local);
				}
				// 外框文件
				if (pic != null) {
					File pic_local = new File(uploadDist + picName);
					pic.transferTo(pic_local);
				}
//				deviceService.saveOrUpdate(device);
				ret.put("code", RespCode.SUCCESS.getCode());

			} catch (Exception e) {
				ret.put("code", RespCode.INTERFACE_FAIL.getCode());
				e.printStackTrace();
			}
		}

//		ret.put("collectionSize", pageDate.getTotal());
//        ret.put("data", vos);
		ret.put("code", RespCode.SUCCESS.getCode());
		return ret;
	}

}
