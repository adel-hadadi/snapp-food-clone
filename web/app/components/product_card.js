import Image from "next/image";
import Link from "next/link";

import { FaShoppingCart, FaStar } from "react-icons/fa";

const ProductCard = ({ product }) => {
    const handleAddToCart = () => {
        const req = {
            product_id: product.id,
            quantity: 1,
        }

        fetch(`http://localhost/api/profile/cart`, {
            method: 'POST',
            body: JSON.stringify(req),
            headers: {
                authorization: 'Bearer ' + localStorage.getItem("access_token")
            }
        })
            .then(res => res.json)
            .then(res => {
                console.log("product added", res.message)
            })
    }

    return (
        <div className="w-60 h-80 border ml-5 rounded-xl shadow overflow-hidden py-5 px-2 relative">
            <section className="w-full h-1/2 relative mb-3 rounded-lg overflow-hidden">
                <Image
                    src={product.image}
                    fill={true}
                    alt={product.name}
                />
            </section>
            <section className="flex flex-col items-center justify-center px-3 h-auto">
                <span className="font-bold mb-1 text-center">{product.name}</span>

                <Link className="mb-3" href={`stores/${product?.store.slug}`}>
                    <span className="text-xs">
                        {product?.store.name}
                    </span>
                </Link>

                <div className="text-xs flex justify-between w-full mb-4">
                    <span>
                        {product.price} تومان
                    </span>
                    <span className="inline-flex items-center">
                        <FaStar className="text-yellow-400 ml-2" /> {"۴.۳"}
                    </span>
                </div>
                <button 
                    className="w-11/12 bg-pink-700 text-sm rounded text-white py-2 px-3 absolute bottom-3 flex justify-between items-center" 
                    onClick={handleAddToCart}
                >
                    <span>
                        افزودن به سبد خرید
                    </span>
                    <FaShoppingCart />
                </button>
            </section>
        </div>

    );
}

export default ProductCard
