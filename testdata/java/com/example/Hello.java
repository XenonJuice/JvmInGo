package com.example;

public class Hello {
    public static void main(String[] args) {
        System.out.println("hello from test class");
    }

    public static int add(int a, int b) {
        return a + b;
    }

    public static int subtract(int a, int b) {
        return a - b;
    }

    public static boolean isPositive(int value) {
        return value > 0;
    }

    public int multiply(int a, int b) {
        return a * b;
    }
}
