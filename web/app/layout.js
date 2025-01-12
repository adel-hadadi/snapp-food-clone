import "./globals.css";

export const metadata = {
  title: "اسنپ فود",
  description: "سفارش آنلاین غذا",
};

export default function RootLayout({ children }) {
  return (
    <html lang="fa" dir="rtl">
      <body className="container mx-auto pt-12 px-32">{children}</body>
    </html>
  );
}
