"use client";

import Image from "next/image";
import Link from "next/link";
import { useEffect, useState } from "react";
import { FaChevronLeft, FaStar } from "react-icons/fa";
import ProductCard from "../components/product_card";
import Store from "@/app/components/store";

const Page = () => {
  const [productCategories, setProductCategories] = useState([]);
  const [stores, setStores] = useState([]);
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const accessToken = "Bearer " + localStorage.getItem("access_token");
    fetch("http://localhost/api/product-categories", {
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .then(({ data }) => {
        setProductCategories(data);
      });

    fetch("http://localhost/api/panel/stores/nearest", {
      headers: {
        "Content-Type": "application/json",
        authorization: accessToken,
      },
    })
      .then((res) => res.json())
      .then(({ data }) => setStores(data))
      .catch((err) => {
        console.log("an error accured => ", err);
      });

    fetch("http://localhost/api/panel/products?sort=rate", {
      headers: {
        "Content-Type": "application/json",
        authorization: accessToken,
      },
    })
      .then((res) => res.json())
      .then(({ data }) => setProducts(data))
      .catch((err) => {
        console.log("an error accured => ", err);
      });
  }, []);

  return (
    <>
      <section className="mb-10">
        <h1 className="font-bold mb-5 text-xl">دسته بندی ها</h1>
        <section>
          {productCategories.length > 0 && (
            <ul className="grid grid-cols-6 gap-6">
              {productCategories.map((c) => (
                <Link
                  key={c.id}
                  href={`panel/product-categories/${c.slug}/stores`}
                  className="relative h-32 bg-white cursor-pointer rounded-lg overflow-hidden border-2 border-white shadow"
                >
                  <Image src={c.image} fill={true} alt={c.name} />
                  <div className="bg-inherit pr-3 pl-4 py-1 absolute bottom-0 right-0 rounded-tl-xl inline-flex items-center">
                    <span className="ml-2">{c.name}</span>
                    <FaChevronLeft />
                  </div>
                </Link>
              ))}
            </ul>
          )}
        </section>
      </section>

      {stores.length > 0 && (
        <section className="mb-10">
          <h1 className="font-bold mb-5 text-xl">نزدیک ترین رستوران ها</h1>
          <div className="flex">
            {stores.map((store) => <Store store={store} />)}
          </div>
        </section>
      )}

      {products.length > 0 && (
        <section className="mb-10">
          <h1 className="font-bold mb-5 text-xl">برترین غذاها</h1>
          <div className="flex">
            {products.map((product) => (
              <ProductCard key={product.id} product={product} />
            ))}
          </div>
        </section>
      )}
    </>
  );
};

export default Page;
