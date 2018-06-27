package com.ngtesting.platform.servlet;

import com.ngtesting.platform.vo.ProgressVo;
import com.sun.javaws.progress.Progress;
import org.apache.commons.fileupload.ProgressListener;
import org.springframework.stereotype.Component;

import javax.servlet.http.HttpSession;

@Component
public class FileUploadProgressListener implements ProgressListener {

    private HttpSession session;

    public void setSession(HttpSession session) {
        this.session = session;
        Progress status = new Progress();//保存上传状态
        session.setAttribute("status", status);
    }

    @Override
    public void update(long bytesRead, long contentLength, int items) {
        ProgressVo status = (ProgressVo) session.getAttribute("status");
        status.setBytesRead(bytesRead);
        status.setContentLength(contentLength);
        status.setItems(items);

    }

}
