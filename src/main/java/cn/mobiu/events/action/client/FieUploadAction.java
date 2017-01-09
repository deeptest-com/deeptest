package cn.mobiu.events.action.client;

import java.io.File;
import java.text.DecimalFormat;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.io.FilenameUtils;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.multipart.MultipartFile;

import com.alibaba.fastjson.JSONObject;

import cn.mobiu.events.constants.Constant;
import cn.mobiu.events.entity.EvtClient;
import cn.mobiu.events.service.ClientService;
import cn.mobiu.events.service.QrcodeService;
import cn.mobiu.events.util.AuthPassport;
import cn.mobiu.events.util.DateUtils;
import cn.mobiu.events.util.FileUtils;

@Controller
@RequestMapping(Constant.API_PATH_CLIENT)
public class FieUploadAction extends BaseAction {

    Log logger = LogFactory.getLog(FieUploadAction.class);

    @Autowired
    private ClientService clientService;
    @Autowired
    private QrcodeService qrcodeService;

    @AuthPassport(validate = true)
    @RequestMapping(value = "sign", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> sign(
            @RequestParam("file") MultipartFile file, HttpServletRequest request, HttpServletResponse response) {
        Map<String, Object> ret = new HashMap<String, Object>();

        String token = request.getParameter("Authorization");
        String eventId = request.getParameter("eventId");
        String extName = request.getParameter("extName");
        
        EvtClient client = clientService.getByToken(token);
        
        extName = extName == null? FilenameUtils.getExtension(file.getOriginalFilename()): extName;
        
        String fileName = UUID.randomUUID().toString() + "." + extName;

        String filePath = saveFile(file, "event/", fileName);

        ret.put("origName", file.getOriginalFilename());
        ret.put("filePath", filePath);
        
        float flt = Float.parseFloat(String.valueOf(file.getSize()));
        String fileSize = new DecimalFormat("##0.00").format(flt / 1000 / 1000);
        ret.put("fileSize", fileSize + 'M');
        
        JSONObject json = qrcodeService.decode(filePath);
        System.out.println("author：" + json.getString("author"));  
        System.out.println("eventId：" + json.getString("eventId"));  
        System.out.println("eventName：" + json.getString("eventName"));
        
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }
    
    @AuthPassport(validate = true)
    @RequestMapping(value = "uploadSingle", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> uploadSingle(
            @RequestParam("file") MultipartFile file, HttpServletRequest request, HttpServletResponse response) {
        Map<String, Object> ret = new HashMap<String, Object>();
        
        EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
        
        String extName = FilenameUtils.getExtension(file.getOriginalFilename());
        String fileName = UUID.randomUUID().toString() + "." + extName;

        String filePath = saveFile(file, "event/", fileName);

        ret.put("origName", file.getOriginalFilename());
        ret.put("filePath", filePath);
        
        float flt = Float.parseFloat(String.valueOf(file.getSize()));
        String fileSize = new DecimalFormat("##0.00").format(flt / 1000 / 1000);
        ret.put("fileSize", fileSize + 'M');

        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

    public String saveFile(MultipartFile file, String uploadRelativeDist, String fileName) {
        String dateDist = DateUtils.getDateNoSeparator();
        
        String localFolder = Constant.GetUploadDir() + uploadRelativeDist + dateDist + "/";
        FileUtils.CreateDirIfNeeded(localFolder);
        
        String localPath = localFolder + fileName;
        File localFile = new File(localPath);
        try {
            file.transferTo(localFile);
        } catch (Exception e) {
            logger.error(e.getStackTrace());
            return null;
        }
        
        return localPath;
    }

}
