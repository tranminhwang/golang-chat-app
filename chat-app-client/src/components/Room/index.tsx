import React from "react";
import { PrimaryButton } from "../shared/Button";

export interface IRoom {
  id: string;
  name: string;
}

const Room: React.FC<IRoom> = ({ id, name }) => {
  return (
    <div className="flex justify-between text-[#0063ec] bg-white rounded-lg px-8 py-6">
      <h3 className="text-2xl font-medium">{name}</h3>
      <button
        type="button"
        className="px-3 py-2 text-sm font-medium text-center text-white bg-blue-700 rounded-full hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
      >
        Join
      </button>
    </div>
  );
};

export default Room;
