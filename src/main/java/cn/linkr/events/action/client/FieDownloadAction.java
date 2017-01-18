package cn.linkr.events.action.client;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

import cn.linkr.events.constants.Constant;

@Controller
@RequestMapping(Constant.API_PATH_CLIENT)
public class FieDownloadAction extends BaseAction {

    Log logger = LogFactory.getLog(FieDownloadAction.class);

    @RequestMapping("/download")
    public String download(String fileName, String currDeviceId, String imgId, HttpServletRequest request,
                           HttpServletResponse response) {
//		SysDevice currDevice = (SysDevice) deviceService.get(SysDevice.class, Long.valueOf(currDeviceId));
//		SysConsumer tstAsset = (SysConsumer) assetService.get(SysConsumer.class, Long.valueOf(imgId));
//		String[] split = fileName.split("\\.");
//		String suffix = split[split.length-1];
//		String reFileName = currDevice.getName().replace(" ", "") + new SimpleDateFormat("yyyy-MM-dd-HH-mm-ss").format(tstAsset.getCreateTime()) + "." + suffix;
//        response.setCharacterEncoding("utf-8");
//        response.setContentType("multipart/form-data");
//        try {
//			response.setHeader("Content-Disposition", "attachment;fileName="
//			        + new String((reFileName).getBytes(), "iso-8859-1"));
//		} catch (UnsupportedEncodingException e1) {
//			e1.printStackTrace();
//		}
//
//        OutputStream os = null;
//        InputStream inputStream = null;
//        try {
//            inputStream = new FileInputStream(
//            		new File(Constant.UPLOAD_ABSOLUTE_DIR + "upload" + fileName));
//
//            os = response.getOutputStream();
//            byte[] b = new byte[2048];
//            int length;
//            while ((length = inputStream.read(b)) > 0) {
//                os.write(b, 0, length);
//            }
//
//             // 这里主要关闭。
//            os.close();
//
//            inputStream.close();
//        } catch (FileNotFoundException e) {
//            e.printStackTrace();
//        } catch (IOException e) {
//            e.printStackTrace();
//        } finally {
//        	try {
//				if (os != null) {
//					os.close();
//				}
//        		if (inputStream != null) {
//        			inputStream.close();
//				}
//			} catch (IOException e) {
//				e.printStackTrace();
//			}
//        }
//            //  返回值要注意，要不然就出现下面这句错误！
//            //java+getOutputStream() has already been called for this response
        return null;
    }
}
