package cn.linkr.events.servlet;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;

import org.apache.commons.fileupload.ProgressListener;

import cn.linkr.events.vo.ProgressEntity;

/**
 * 文件上传进度监听服务
 *
 * @author xuxiang
 * @version $Id$
 * @see
 */
public class PJProgressListener implements ProgressListener {
    private HttpServletRequest request;
    private HttpSession session;

    public PJProgressListener() {
    }

    public PJProgressListener(HttpServletRequest _request) {
        request = _request;
        session = request.getSession();

        String fileversion = request.getParameter("fileversion");
        String progressSessionId = String.format("%s_upload_ps", fileversion);

        ProgressEntity progressEntity = new ProgressEntity();
        session.setAttribute(progressSessionId, progressEntity);
    }

    public void update(long pBytesRead, long pContentLength, int pItems) {
        String fileversion = request.getParameter("fileversion");
        String progressSessionId = String.format("%s_upload_ps", fileversion);

        ProgressEntity ps = (ProgressEntity) session.getAttribute(progressSessionId);
        ps.setpBytesRead(pBytesRead);
        ps.setpContentLength(pContentLength);
        ps.setpItems(pItems);
        session.setAttribute(progressSessionId, ps);
    }

}
