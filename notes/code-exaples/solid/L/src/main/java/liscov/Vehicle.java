package liscov;

public class Vehicle {
    protected String type;
    protected int age;

    public Vehicle(String type, int age) {
        this.type = type;
        this.age = age;
    }

    public void slowDown() {
        System.out.println("Vehicle.enclosing_method()");
    }

    public void speedUp() {
        System.out.println("Vehicle.enclosing_method()");
    }

    public void fuel() {
        System.out.println("Vehicle.fuel()");
    }

}
