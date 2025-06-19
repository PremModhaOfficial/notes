package liscov;

/**
 * Hello world!
 *
 */
public class App {
    public static void main(String[] args) {
        System.out.println("App.main()");

        Vehicle vehiclePet = new PetrolCar("Innova", 4);
        Vehicle vehicleEle = new ElectricCar("Tesla", 1);

        vehiclePet.fuel();
        vehicleEle.fuel(); // errors
    }
}
