/**
  * Copyright 2017 bejson.com
  */
package com.ngtesting.platform.vo;

public class AiRunMlf {

    private String category;
    private String file;
    private String name;
    private String path;
    private String regexInput;

    public String getRegexInput() {
        return regexInput;
    }

    public void setRegexInput(String regexInput) {
        this.regexInput = regexInput;
    }

    public void setCategory(String category) {
         this.category = category;
     }
     public String getCategory() {
         return category;
     }

    public String getFile() {
        return file;
    }

    public void setFile(String file) {
        this.file = file;
    }

    public String getPath() {
        return path;
    }

    public void setPath(String path) {
        this.path = path;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
