package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.ConstantIssue;
import com.ngtesting.platform.dao.IssueCustomFieldDao;
import com.ngtesting.platform.dao.IssueCustomFieldOptionDao;
import com.ngtesting.platform.dao.IssuePageElementDao;
import com.ngtesting.platform.model.IsuCustomField;
import com.ngtesting.platform.service.intf.IssueCustomFieldService;
import com.ngtesting.platform.service.intf.ProjectService;
import com.ngtesting.platform.service.intf.TestCustomFieldOptionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.LinkedList;
import java.util.List;

@Service
public class IssueCustomFieldServiceImpl extends BaseServiceImpl implements IssueCustomFieldService {
    @Autowired
    IssueCustomFieldDao customFieldDao;
    @Autowired
    IssueCustomFieldOptionDao customFieldOptionDao;

    @Autowired
    IssuePageElementDao elementDao;

    @Autowired
    ProjectService projectService;
    @Autowired
    TestCustomFieldOptionService customFieldOptionService;

    @Override
    public List<IsuCustomField> list(Integer orgId) {
        List<IsuCustomField> ls = customFieldDao.list(orgId);

        return ls;
    }

    @Override
    public IsuCustomField get(Integer id, Integer orgId) {
        return customFieldDao.getDetail(id, orgId);
    }

    @Override
    @Transactional
    public IsuCustomField save(IsuCustomField vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            Integer maxOrder = customFieldDao.getMaxOrdrNumb(orgId);
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            customFieldDao.save(vo);
//            if (vo.getInput().equals(ConstantIssue.IssueInput.dropdown)
//                    || vo.getInput().equals(ConstantIssue.IssueInput.radio)
//                    || vo.getInput().equals(ConstantIssue.IssueInput.checkbox)) {
//                customFieldOptionDao.saveAll(vo.getId(), vo.getOptions());
//            }
        } else {
            Integer count = customFieldDao.update(vo);
            if (count == 0) {
                return null;
            }

            elementDao.updateFromCustomField(vo);
        }

        return vo;
    }

    @Override
    public Boolean delete(Integer id, Integer orgId) {
        Integer count = customFieldDao.delete(id, orgId);
        if (count == 0) {
            return false;
        }

        return true;
    }

    @Override
    public Boolean changeOrderPers(Integer id, String act, Integer orgId) {
        IsuCustomField curr = customFieldDao.get(id, orgId);
        if (curr == null) {
            return false;
        }

        IsuCustomField neighbor = null;
        if ("up".equals(act)) {
            neighbor = customFieldDao.getPrev(curr.getOrdr(), orgId);
        } else if ("down".equals(act)) {
            neighbor = customFieldDao.getNext(curr.getOrdr(), orgId);
        }
        if (neighbor == null) {
            return false;
        }

        Integer currOrder = curr.getOrdr();
        Integer neighborOrder = neighbor.getOrdr();
        customFieldDao.setOrder(id, neighborOrder, orgId);
        customFieldDao.setOrder(neighbor.getId(), currOrder, orgId);

        return true;
    }

    @Override
    public String getLastUnusedColumn(Integer orgId) {
        List<String> ls = customFieldDao.getLastUnusedColumn(orgId);

        String ret = null;
        for (int i = 1; i <= 20; i++) {
            String prop = "prop" + String.format("%02d", i);
            if (!ls.contains(prop)) {
                ret = prop;
                break;
            }
        }

        return ret;
    }

    @Override
    public List<String> listType() {
        List<String> ls = new LinkedList<String>();
        for (ConstantIssue.IssueType item : ConstantIssue.IssueType.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

    @Override
    public List<String> listInput() {
        List<String> ls = new LinkedList<String>();
        for (ConstantIssue.IssueInput item : ConstantIssue.IssueInput.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

    @Override
    public List<String> listFormat() {
        List<String> ls = new LinkedList();
        for (ConstantIssue.TextFormat item : ConstantIssue.TextFormat.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

}
