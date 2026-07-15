# ⚡ TechMart — Full-Stack E-Commerce Application

TechMart is a full-stack e-commerce application for selling electronic products, built from scratch using **Golang** (backend), **Angular** (frontend), and **MySQL** (database). This project covers a complete e-commerce business flow — from authentication, product management, shopping cart, checkout, manual payment, to an admin dashboard with sales statistics.

## 🔗 Live Demo

- **Frontend:** _(to be added after deployment)_
- **Backend API:** _(to be added after deployment)_

## 🛠️ Tech Stack

**Backend**
- Golang + [Gin](https://gin-gonic.com/) (web framework)
- [GORM](https://gorm.io/) (ORM)
- MySQL (database)
- JWT (authentication)
- Bcrypt (password hashing)

**Frontend**
- Angular 18 (standalone components)
- Custom CSS with a minimalist design system
- Chart.js (data visualization)
- RxJS

**Others**
- RESTful API
- Role-Based Access Control (Admin & Customer)
- File upload (product images, profile photos, payment proofs)

## ✨ Key Features

### Authentication & Security
- Register & Login with JWT
- Role-based access control (Admin vs Customer)
- Frontend route guards (protecting pages based on role)
- Forgot password with security question
- Change password from profile page

### For Customers
- Browse product catalog with search, category filter, and pagination
- Product detail page
- Shopping cart (add, update quantity, remove items)
- Checkout & manual payment (upload transfer proof)
- Order history with status tracking
- Edit profile (name, gender, profile photo)
- In-app notifications (order status updates, etc.)

### For Admins
- Dashboard with store statistics (revenue chart, best-selling products, order status breakdown)
- Product management (full CRUD + image upload)
- Category management
- Order management (payment verification, status updates)
- Customer management (view profile & purchase history of each customer)
- Admins cannot make purchases (consistent role separation)

## 🗂️ Project Structure

```
techmart-ecommerce/
├── backend/              # Golang REST API
│   ├── config/           # Database configuration
│   ├── controllers/      # Business logic
│   ├── middleware/       # Auth & role middleware
│   ├── models/           # Data structures (GORM models)
│   ├── routes/           # API route definitions
│   ├── utils/            # Helper functions (JWT, hashing, etc.)
│   └── main.go
└── frontend/             # Angular application
    └── src/app/
        ├── auth/          # Login, Register, Forgot Password
        ├── dashboard/     # Main dashboard
        ├── products/      # Catalog, Detail, Product Management
        ├── cart/          # Shopping cart
        ├── orders/        # Order history & Order Management
        ├── profile/       # Edit profile
        ├── users/         # Customer Management (admin)
        ├── stats/         # Statistics Dashboard (admin)
        └── notifications/ # In-app notifications
```

## 🚀 Running Locally

### Prerequisites
- [Go](https://go.dev/dl/) 1.22+
- [Node.js](https://nodejs.org/) 18+ & Angular CLI
- MySQL

### 1. Clone the Repository
```bash
git clone https://github.com/Bronsonsiallagan/techmart-ecommerce.git
cd techmart-ecommerce
```

### 2. Backend Setup
```bash
cd backend
go mod tidy
```

Create a `.env` file (see `.env.example` for reference):
```env
DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=techmart
JWT_SECRET=replace_with_your_own_secret_string
```

Create the MySQL database:
```sql
CREATE DATABASE techmart;
```

Run the server:
```bash
go run main.go
```
The backend will run on `http://localhost:8080`

### 3. Frontend Setup
Open a new terminal:
```bash
cd frontend
npm install
ng serve
```
The frontend will run on `http://localhost:4200`

## 📋 API Endpoints (Summary)

| Method | Endpoint | Description | Access |
|--------|----------|-------------|--------|
| POST | `/api/auth/register` | Register a new account | Public |
| POST | `/api/auth/login` | Login | Public |
| GET | `/api/products` | List all products | Public |
| POST | `/api/products` | Create a product | Admin |
| GET | `/api/cart` | View shopping cart | Customer |
| POST | `/api/orders` | Checkout | Customer |
| GET | `/api/admin/orders` | View all orders | Admin |
| GET | `/api/admin/stats` | Store statistics | Admin |

_Full API documentation can be found in `routes/routes.go`._

## 🎯 Technical Highlights

- **Role-Based Access Control** enforced consistently on both backend (middleware) and frontend (route guards), not just hidden UI elements
- **Data integrity protection** — products that have already been ordered cannot be deleted, preserving transaction history
- **Pagination** implemented at the backend level for better performance on larger datasets
- **File upload** support for product images, profile photos, and payment proofs
- **In-app notification system** with polling for near real-time updates

## 👤 Contact

**Bronson Siallagan**
GitHub: [@Bronsonsiallagan](https://github.com/Bronsonsiallagan)

---

_This project was built as part of a full-stack web development portfolio._
