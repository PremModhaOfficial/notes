package liscov;

public class ElectricCar extends Vehicle {

    public ElectricCar(String type, int age) {
        super(type, age);
    }

    @Override
    public void fuel() {
        throw new Error("what no fuel man");
    }

    @Override
    public void slowDown() {
        System.out.println("ElectricCar.enclosing_method()");
    }

    @Override
    public void speedUp() {
        System.out.println("ElectricCar.speedUp()");
    }

}
