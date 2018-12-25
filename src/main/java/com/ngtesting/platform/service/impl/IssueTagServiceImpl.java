package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.dao.IssueTagDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.IsuTag;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueTagService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class IssueTagServiceImpl extends BaseServiceImpl implements IssueTagService {
    @Autowired
    IssueTagDao issueTagDao;

    @Autowired
    IssueDao issueDao;

    @Override
    public List<IsuTag> search(Integer issueId, Integer orgId, String keywords, List<Integer> exceptIds) {
        return issueTagDao.search(issueId, orgId, keywords, exceptIds);
    }

    @Override
    @Transactional
    public Boolean save(Integer issueId, List<Map> tags, TstUser user) {
        IsuIssue issue = issueDao.get(issueId, user.getId(), user.getDefaultPrjId());
        if (issue == null) {
            return false;
        }

        List<IsuTag> newTags = new LinkedList<>();
        List<IsuTag> allTags = new LinkedList<>();
        String tagStr = "";

        for (Map tag : tags) {
            IsuTag t = new IsuTag(tag.get("name").toString(), user.getDefaultOrgId(), user.getId());
            if (tag.get("id") == null) {
                newTags.add(t);
            } else {
                t.setId(Integer.valueOf(tag.get("id").toString()));
            }

            allTags.add(t);

            if (!"".equals(tagStr)) {
                tagStr += ",";
            }
            tagStr += t.getName();
        }

        if (newTags.size() > 0) {
            issueTagDao.save(newTags);
        }

        issueTagDao.removeRelations(issueId);
        issueTagDao.saveRelations(issueId, allTags);

        issueTagDao.updateTagField(issueId, tagStr);

        return true;
    }

}

