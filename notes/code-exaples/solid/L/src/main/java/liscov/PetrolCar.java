package liscov;

public class PetrolCar extends Vehicle {

    public PetrolCar(String type, int age) {
        super(type, age);
    }

    @Override
    public void slowDown() {
        System.out.println("PetrolCar.slowDown()");
    }

    @Override
    public void speedUp() {
        System.out.println("PetrolCar.speedUp()");
    }

    @Override
    public void fuel() {
        System.out.println("PetrolCar.fuel()");
    }

}
