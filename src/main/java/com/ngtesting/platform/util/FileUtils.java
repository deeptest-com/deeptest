package com.ngtesting.platform.util;

import java.awt.image.BufferedImage;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.OutputStream;
import java.util.Scanner;

import javax.imageio.ImageIO;

import org.apache.commons.io.FilenameUtils;
import org.springframework.web.multipart.MultipartFile;

public class FileUtils {
	
	public static final int THUMB_WIDTH = 100;
	public static final int THUMB_HEIGHT = 100;

    public static String SaveFile(MultipartFile file, String uploadRelativeDist, String fileName) {
        String dateDist = DateUtils.getDateNoSeparator();
        
        String uploadPath = Constant.FTP_UPLOAD_DIR + uploadRelativeDist + dateDist + "/";
        String localFolder = Constant.WORK_DIR + uploadPath;
        
        FileUtils.CreateDirIfNeeded(localFolder);
        
        String localPath = localFolder + fileName;
        File localFile = new File(localPath);
        try {
            file.transferTo(localFile);
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }
        
        FileUtils.Thumb(localPath);
        return uploadPath + fileName;
    }
    
    public static String Thumb(String srcPath){ 
        String imageType = FilenameUtils.getExtension(srcPath);
        String des  = srcPath.replace(".", "-thumb."); 
        
        OutputStream os = null;  
        try{
        	File file = new File(srcPath);
        	BufferedImage image = ImageIO.read(file);  
            int width = image.getWidth(null);//原图宽度  
            int height = image.getHeight(null);//原图高度  
              
            int rate1 = width / THUMB_WIDTH;//宽度缩略比例  
            int rate2 = height / THUMB_HEIGHT;//高度缩略比例  
              
            int rate = 0;  
            if(rate1 > rate2){//宽度缩略比例大于高度缩略比例，使用宽度缩略比例  
                rate = rate1;  
            }else{  
                rate = rate2;  
            }
            if (rate == 0) {
            	rate = 1;
            }
              
            //计算缩略图最终的宽度和高度  
            int newWidth = width / rate;  
            int newHeight = height / rate;  
              
            BufferedImage bufferedImage = new BufferedImage(newWidth, newHeight, BufferedImage.TYPE_INT_RGB);  
            bufferedImage.getGraphics().drawImage(image.getScaledInstance(newWidth, newHeight, image.SCALE_SMOOTH),  
                    0,0,null);  
 
            os = new FileOutputStream(des);
            
            ImageIO.write(bufferedImage, imageType, os);
        }catch(Exception e){  
            e.printStackTrace();  
        }finally{  
            if(os!=null){  
                try {  
                    os.close();  
                } catch (IOException e) {
                    e.printStackTrace();  
                }  
            }  
        }  
        return des;  
    }
    
    public static String ReadFile(String fullFileName) {
        File file = new File(fullFileName);
        Scanner scanner = null;
        StringBuilder buffer = new StringBuilder();
        try {
            scanner = new Scanner(file, "utf-8");
            while (scanner.hasNextLine()) {
                buffer.append(scanner.nextLine());
            }
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } finally {
            if (scanner != null) {
                scanner.close();
            }
        }

        return buffer.toString();
    }
    
    public static void CreateDirIfNeeded(String dir) {
	  File file = new File(dir);
	  if(!file.exists()){
		  file.mkdirs();
	  }
    }
    
}