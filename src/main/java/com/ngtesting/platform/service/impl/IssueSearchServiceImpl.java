package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueSearchDao;
import com.ngtesting.platform.service.intf.IssueSearchService;
import com.ngtesting.platform.utils.StringUtil;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class IssueSearchServiceImpl extends BaseServiceImpl implements IssueSearchService {
    Log logger = LogFactory.getLog(IssueSearchServiceImpl.class);

    @Autowired
    IssueSearchDao issueSearchDao;

    @Override
	public List<Map> idAndTitleSearch(String text, List<Integer> exceptIds, Integer projectId) {
        List<Map> ls = new LinkedList<>();

        if (text.startsWith("IS-")) {
            text = text.split("-")[1];

            if (StringUtil.isNumeric(text)) {
                ls = issueSearchDao.idSearch(text, exceptIds, projectId);
            }
        } else {
            ls = issueSearchDao.titleSearch(text, exceptIds, projectId);
        }

		return ls;
	}

}

