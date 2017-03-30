package com.ngtesting.platform.dao.impl;

import org.hibernate.SessionFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import com.ngtesting.platform.dao.ProjectDao;

@Repository("projectDao")
@SuppressWarnings("all")
public class ProjectDaoImpl implements ProjectDao {

    @Autowired
    private SessionFactory sessionFactory;

}
