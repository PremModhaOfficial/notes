
public class EmployeeAfter {
	private String name;
	private String address;

	public EmployeeAfter(String name, String address) {
		this.setName(name);
		this.setAddress(address);
	}

	public String getAddress() {
		return address;

	}

	public void setAddress(String address) {
		this.address = address;

	}

	public String getName() {
		return name;

	}

	public void setName(String name) {
		this.name = name;

	}

	// public void saveToDatabase() {
	// System.out.println("saved to db: " + this.name);
	// System.out.println("saved to db: " + this.address);
	// }
	//
	// public void printReport() {
	// System.out.println("report: " + this.name);
	// System.out.println("report: " + this.address);
	// }

}

class EmployeeSaveToDB {
	public void saveToDB(EmployeeAfter e) {
		System.out.println("saved to db: " + e.getName());
		System.out.println("saved to db: " + e.getAddress());
	}
}

class EmployeeReport {
	public void createReport(EmployeeAfter e) {
		System.out.println("repot" + e.getName());
		System.out.println("repot" + e.getAddress());
	}
}
