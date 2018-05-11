package com.ngtesting.platform.action;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.OrgService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.DuiUtils;
import com.ngtesting.platform.util.FileUtils;
import com.ngtesting.platform.vo.UserVo;
import org.apache.commons.io.FilenameUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletRequest;
import java.util.*;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "aiutils/")
public class AiUtilsAction extends BaseAction {
    @Autowired
    OrgService orgService;

	@AuthPassport(validate=false)
	@RequestMapping(value = "hresults", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> hresults(@RequestParam("files") List<MultipartFile> files,
										HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();

		List<String> paths = new LinkedList<>();
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		for (MultipartFile file : files) {
            String origName = file.getOriginalFilename();
		    String extName = FilenameUtils.getExtension(origName);
		    String fileName = UUID.randomUUID().toString() + "-"
                    + origName.split("\\.")[0] + "." + extName;

            String uploadPath = FileUtils.SaveFile(file, "data/", fileName, false);
            paths.add(uploadPath);
        }

        String outputPath = FileUtils.RandomFilePath(paths.get(1),"txt").replace(".txt", "-result.txt");
        String cmdOutput = DuiUtils.HResultsCall(Constant.WORK_DIR + paths.get(0),
                Constant.WORK_DIR + paths.get(1), Constant.WORK_DIR + outputPath);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("data", outputPath);
        ret.put("cmdOutput", cmdOutput);

		return ret;
	}

}
