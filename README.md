# Product Sorting Service

This project provides a sorting service for products based on different criteria, such as **price, sales-to-views ratio, and creation date**. It is implemented in **Go** with sorting strategies that can be extended easily.

## Features

- **Sort by Price** (Descending - highest price first)
- **Sort by Sales-to-Views Ratio** (Higher conversion rate first)
- **Sort by Creation Date** (Newest first)
- Uses a flexible **sorting manager** to register and apply different sorting strategies.

## Getting Started

### **1. Install Go**

Ensure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).

### **2. Clone the Repository**

```sh
git clone https://github.com/yourusername/your-repo-name.git
cd your-repo-name
```

### **3. Build the Project**

```sh
make build
```

### **4. Run the Program**

```sh
make run
```

### **5. Run Tests**

```sh
make test
```

### **6. Format Code**

```sh
make fmt
```

### **7. List Available Sorters**

```sh
make sorters
```

## Sorting Methods

The following sorting methods are available:

| Sorter Name      | Description                                        |
| ---------------- | -------------------------------------------------- |
| `Price`          | Sorts products by price (highest to lowest).       |
| `SalesViewRatio` | Sorts by the highest sales-to-views ratio.         |
| `CreationDate`   | Sorts by the most recently created products first. |

## Project Structure

```
/your-repo-name
│── main.go                  # Main entry point
│── products.go              # Product struct and related functions
│── sorters.go               # Sorting strategies (Price, SalesViewRatio, CreationDate)
│── manager.go               # Sorter manager implementation
│── products_test.go         # Unit tests for sorting logic
│── Makefile                 # Build and automation commands
│── README.md                # Documentation
```

## License

This project is open-source. You are free to modify and distribute it under the **MIT License**.
