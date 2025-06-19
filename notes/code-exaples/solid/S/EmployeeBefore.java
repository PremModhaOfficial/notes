
public class EmployeeBefore {
	private String name;
	private String address;

	public EmployeeBefore(String name, String address) {
		this.name = name;
		this.address = address;
	}

	public void saveToDatabase() {
		System.out.println("saved to db: " + this.name);
		System.out.println("saved to db: " + this.address);
	}

	public void printReport() {
		System.out.println("report: " + this.name);
		System.out.println("report: " + this.address);
	}

}
