const Modal = ({ isOpen, handleToggleModal, children, title }) => {
    if (isOpen) {
        return (
            <div
                className="fixed inset-0 flex
                        items-center justify-center
                        bg-black bg-opacity-50 text-right z-50"
            >
                <div
                    className="bg-white rounded-lg
                            shadow-lg p-6 max-w-md
                            w-full relative"
                >
                    <button
                        className="absolute top-2 right-2
                               text-gray-500 hover:text-gray-700"
                        onClick={handleToggleModal}
                    >
                        &#x2715; {/* Close button */}
                    </button>
                    <h1 className="text-center mb-4">{title}</h1>
                    <section>{children}</section>
                </div>
            </div>
        );
    }
};

export default Modal;
