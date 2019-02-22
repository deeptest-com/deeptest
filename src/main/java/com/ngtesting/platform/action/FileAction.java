package com.ngtesting.platform.action;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.intf.UserService;
import com.ngtesting.platform.utils.FileUtil;
import org.apache.commons.io.FileUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletRequest;
import java.io.File;
import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.text.DecimalFormat;
import java.util.HashMap;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH + "/")
public class FileAction extends BaseAction {
    @Autowired
    UserService userService;

    @ResponseBody
    @PostMapping("/uploadSingle")
    public Map<String, Object> uploadSingle(
            @RequestParam("file") MultipartFile file, HttpServletRequest request) {
        Map<String, Object> ret = new HashMap<String, Object>();

        String origName = file.getOriginalFilename();
        String fileName = FileUtil.UuidFileName(origName);

        String uploadPath = FileUtil.SaveFile(file, "data/", fileName);

        ret.put("origName", origName);
        ret.put("uploadPath", uploadPath);

        float flt = Float.parseFloat(String.valueOf(file.getSize()));
        String fileSize = new DecimalFormat("##0.00").format(flt / 1000 / 1000);
        ret.put("fileSize", fileSize + 'M');

        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

    @GetMapping("/download")
    @RequestMapping(value = "download", method = RequestMethod.GET)
    ResponseEntity<byte[]> download(HttpServletRequest request) {
        String path = request.getParameter("path");
        String name = request.getParameter("name");

        //设置http协议头部
        HttpHeaders headers = new HttpHeaders();

        //设置文件名
        String fileName = "";
        try {
            fileName = new String(name.getBytes("UTF-8"), "iso-8859-1"); //为了解决中文名称乱码问题
        } catch (UnsupportedEncodingException e) {
            e.printStackTrace();
        }

        //头部设置文件类型
        headers.setContentDispositionFormData("attachment", fileName);
        headers.setContentType(MediaType.APPLICATION_OCTET_STREAM);

        //获取文件路径
        String targetDirectory = Constant.WORK_DIR + path;
        File file = new File(targetDirectory);

        //返回文件字节数组
        try {
            return new ResponseEntity<byte[]>(FileUtils.readFileToByteArray(file), headers, HttpStatus.CREATED);
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }

}
