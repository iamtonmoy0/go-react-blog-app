import { Outlet } from "react-router-dom";
import Header from "./components/Header";

function App() {
  return (
    <>
      <div className="flex col-3">
        <Header />
        <div className="w-full">
          <Outlet />
        </div>
      </div>
    </>
  );
}

export default App;
