package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IsuTqlDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVer;
import com.ngtesting.platform.service.IsuTqlService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
import java.net.URLEncoder;
import java.util.List;

@Service
public class IsuTqlServiceImpl extends BaseServiceImpl implements IsuTqlService {
    UserDao userDao;
    @Autowired
    IsuTqlDao isuTqlDao;

    @Override
    public List<TstVer> getFilters(String tql) {
        try {
            tql = URLEncoder.encode("project=350 " + // sample for testing
                    "AND status='in_progress' " +
                    "AND prop01='val' ORDER BY status ASC", "UTF-8");

            tql =  URLDecoder.decode(tql, "UTF-8");
        } catch (UnsupportedEncodingException e) {
            e.printStackTrace();
        }


        return null;
    }

    @Override
    public Boolean save(Integer caseId, String name, String path, TstUser user) {
        return null;
    }

    @Override
    public Boolean delete(Integer id, TstUser user) {
        return null;
    }
}
