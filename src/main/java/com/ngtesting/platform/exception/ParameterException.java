package com.ngtesting.platform.exception;

public class ParameterException extends RuntimeException {

    /**
     * @Fields serialVersionUID
     */
    private static final long serialVersionUID = 1L;

    public ParameterException() {
        super();
    }

    public ParameterException(String message) {
        super(message);
    }

    public ParameterException(Throwable cause) {
        super(cause);
    }

    public ParameterException(String message, Throwable cause) {
        super(message, cause);
    }

}
