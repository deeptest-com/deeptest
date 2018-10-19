package com.ngtesting.platform.utils;

import com.ngtesting.platform.config.Constant;
import org.apache.commons.io.FilenameUtils;
import org.springframework.web.multipart.MultipartFile;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.*;
import java.util.Scanner;
import java.util.UUID;

public class FileUtil {

	public static final int THUMB_WIDTH = 100;
	public static final int THUMB_HEIGHT = 100;

    public static String SaveFile(MultipartFile file, String uploadRelativeDist, String fileName) {
        return SaveFile(file, uploadRelativeDist, fileName, false);
    }
    public static String SaveFile(MultipartFile file, String uploadRelativeDist, String fileName, boolean thumb) {
        String dateDist = DateUtil.GetDateNoSeparator();

        String uploadPath = Constant.FTP_UPLOAD_DIR + uploadRelativeDist + dateDist + "/";
        String localFolder = Constant.WORK_DIR + uploadPath;

        FileUtil.CreateDirIfNeeded(localFolder);

        String localPath = localFolder + fileName;
        File localFile = new File(localPath);
        try {
            file.transferTo(localFile);
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }

        if (thumb) {
            FileUtil.Thumb(localPath);
        }

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

    public static String RandomFilePath(String brother, String ext) {
        String dir = brother.substring(0, brother.lastIndexOf("/") + 1) ;
        String fileName = UUID.randomUUID().toString() + "." + ext;
        String path = dir + fileName;

        return path;
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

//    public static List<AiRunMlf> ListMlf(String dir, String testType) {
//        List<AiRunMlf> mlfs = new LinkedList<>();
//        FileUtil.TraverseFolder(dir, mlfs, testType);
//
//        return mlfs;
//    }

//    public static void TraverseFolder(String path, List<AiRunMlf> mlfs, String testType) {
//        File file = new File(path);
//        if (file.exists()) {
//            File[] files = file.listFiles();
//            if (files.length == 0) {
//                return;
//            } else {
//                for (File file2 : files) {
//                    if (file2.isDirectory()) {
//                        if (file2.getAbsolutePath().indexOf("__MACOSX") == -1) {
//                            TraverseFolder(file2.getAbsolutePath(), mlfs, testType);
//                        }
//                    } else {
//                        String fileName = file2.getAbsolutePath();
//
//                        System.out.println("文件:" + file2.getAbsolutePath());
//
//                        String mlfPath = null;
//                        if ("nlu-sent".equals(testType) && fileName.indexOf(".txt") > -1) {
//                            mlfPath = file2.getAbsolutePath();
//                        } else if (!"nlu-sent".equals(testType) && fileName.indexOf(".mlf") > -1) {
//                            mlfPath = file2.getAbsolutePath();
//                        }
//                        if (mlfPath != null) {
//                            mlfPath = "work/" + mlfPath.split("work/")[1];
//
//                            AiRunMlf mlf = new AiRunMlf();
//                            mlf.setPath(mlfPath);
//                            mlfs.add(mlf);
//                        }
//                    }
//                }
//            }
//        } else {
//            System.out.println("文件不存在!");
//        }
//    }

}
