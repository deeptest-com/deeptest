package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.dao.IssuePageDao;
import com.ngtesting.platform.dao.IssuePageElementDao;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.IssueCommentsService;
import com.ngtesting.platform.service.intf.IssueHistoryService;
import com.ngtesting.platform.service.intf.IssueService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.*;

@Service
public class IssueServiceImpl extends BaseServiceImpl implements IssueService {
    Log logger = LogFactory.getLog(IssueServiceImpl.class);

    @Autowired
    IssueHistoryService issueHistoryService;

    @Autowired
    IssueDao issueDao;

	@Autowired
	IssuePageDao pageDao;
    @Autowired
    IssuePageElementDao pageElementDao;
    @Autowired
    IssueCommentsService issueCommentsService;

    @Override
	public IsuIssue get(Integer id, Integer userId, Integer prjId) {
		IsuIssue po = issueDao.get(id, userId, prjId);

		return po;
	}

    @Override
    public IsuIssue getDetail(Integer id, Integer userId, Integer prjId) {
        IsuIssue po = issueDao.getDetail(id, userId, prjId);

        return po;
    }

    @Override
    public IsuIssue getData(Integer id, Integer userId, Integer prjId) {
        IsuIssue po = issueDao.getData(id, userId, prjId);

        return po;
    }

    @Override
	public IsuPage getPage(Integer orgId, Integer prjId, String opt) {
        IsuType type = getProjectDefaultType(orgId, prjId);
        Map<String, Integer> pageMap = getProjectDefaultPages(orgId, prjId, type.getId());

		IsuPage page = pageDao.get(pageMap.get(opt), orgId);
		return page;
	}

    @Override
    @Transactional
    public IsuIssue save(JSONObject issue, Integer pageId, TstUser user) {
        List<IsuPageElement> elemObjs = pageElementDao.listElementByPageId(pageId);
        List<IsuPageElement> elems = genElems(elemObjs);

        String uuid = UUID.randomUUID().toString().replace("-", "");
        issue.put("uuid", uuid);

        List<Object> params = genParams(issue, elems, true);

        List<IsuPageElement> elems1 = new LinkedList<>();
        List<Object> params1 = new LinkedList<>();
        List<IsuPageElement> elems2 = new LinkedList<>();
        List<Object> params2 = new LinkedList<>();

        genDataForExtTable(elems, params, elems1, params1, elems2, params2);

        Integer count = issueDao.save(elems1, params1);
        IsuIssue po = null;
        if (count > 0) {
            po = issueDao.getByUuid(uuid);
            count = issueDao.saveExt(elems2, params2, po.getId());

            po.setCreatorId(user.getId());
            po.setOrgId(user.getDefaultOrgId());
            po.setProjectId(user.getDefaultPrjId());
            issueDao.setDefaultVal(po);
        }

        issueHistoryService.saveHistory(user, Constant.EntityAct.create, po.getId(),null);

        if (count > 0) {
            po = issueDao.get(po.getId(), user.getId(), user.getDefaultPrjId());
        }

        return po;
    }

    @Override
    @Transactional
    public Boolean update(JSONObject issue, Integer pageId, TstUser user) {
        List<IsuPageElement> elems = pageElementDao.listElementByPageId(pageId);

        IsuIssue po = null;
        List<Object> params = genParams(issue, elems, false);

        List<IsuPageElement> elems1 = new LinkedList<>();
        List<Object> params1 = new LinkedList<>();
        List<IsuPageElement> elems2 = new LinkedList<>();
        List<Object> params2 = new LinkedList<>();

        genDataForExtTable(elems, params, elems1, params1, elems2, params2);

        Integer count = issueDao.update(elems1, params1, issue.getInteger("id"), user.getDefaultOrgId());
        if (count > 0 && params2.size() > 0) {
            count = issueDao.updateExt(elems2, params2, issue.getInteger("id"));
        }

        issueHistoryService.saveHistory(user, Constant.EntityAct.update, issue.getInteger("id"),null);

        return count > 0;
    }

    @Override
    @Transactional
    public IsuIssue updateField(JSONObject json, TstUser user) {
        Integer projectId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");
        String code = json.getString("code");
        String value = json.getString("value");
        String label = json.getString("label");

        Integer count;
        if (!code.startsWith("prop")) {
            count = issueDao.updateProp(id, code, value, projectId);
        } else {
            count = issueDao.updatePropExt(id, code, value);
        }
        if (count == 0) {
            return null;
        }

        issueHistoryService.saveHistory(user, Constant.EntityAct.update, id, label);
        IsuIssue po = issueDao.get(id, user.getId(), user.getDefaultPrjId());

        return po;
    }

    private List<IsuPageElement> genElems(List<IsuPageElement> elemObjs) {
        List<IsuPageElement> elems = new LinkedList<>();
        for (IsuPageElement elem : elemObjs) {
            elems.add(elem);
        }

        elems.add(new IsuPageElement("uuid", ""));

        return elems;
    }

    private List genParams(JSONObject issue, List<IsuPageElement> elems, Boolean isSave) {
        List<Object> params = new LinkedList<>();

        int i = 0;
        for (IsuPageElement elem : elems) {
            String code = elem.getColCode();

            switch(elem.getInput()){
                case "date":
                    params.add(issue.get(code)!=null?issue.getDate(code): null);
                    break;

                default:
                    params.add(issue.get(code));
                    break;
            }
        }

        if (isSave) {
            params.add(issue.getString("uuid"));
        }

        return params;
    }

    @Override
    public IsuType getProjectDefaultType(Integer orgId, Integer prjId) {
        IsuType type = issueDao.getProjectDefaultType(orgId, prjId);
        return type;
    }

    @Override
    public Map<String, Integer>  getProjectDefaultPages(Integer orgId, Integer prjId, Integer typeId) {
        List<Map<String, Object>> ls = issueDao.getProjectDefaultPages(orgId, prjId, typeId);

        Map<String, Integer> ret = new HashMap();
        for (Map<String, Object> map : ls) {
            ret.put(map.get("opt").toString(), Integer.valueOf(map.get("pageId").toString()));
        }
        return ret;
    }

    @Override
    public void genDataForExtTable(List<IsuPageElement> elems,
                                   List<Object> params,
                                   List<IsuPageElement> elems1,
                                   List<Object> params1,
                                   List<IsuPageElement> elems2,
                                   List<Object> params2) {
        int i = 0;
        for (IsuPageElement elem: elems) {
            if (!elem.getColCode().startsWith("prop")) {
                elems1.add(elem);
                params1.add(params.get(i));
            } else {
                elems2.add(elem);
                params2.add(params.get(i));
            }

            i++;
        }
    }

    @Override
    public void delete(Integer id, TstUser user) {
        issueDao.delete(id, user.getDefaultPrjId());
    }

}

