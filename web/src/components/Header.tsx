import { useNavigate } from "react-router-dom";

interface HeaderProps {
  className?: string;
  style?: React.CSSProperties;
}

const Header = (props: HeaderProps) => {
  const navigate = useNavigate();
  return (
    <div
      className={`w-full items-center justify-start border-b border-neutral-800 transition-all duration-200 lg:pl-32 lg:pr-32 ${props.className}`}
      style={{ ...props.style }}
    >
      <div className="flex flex-row items-center justify-between">
        <div
          className="flex cursor-pointer flex-row items-center p-4"
          onClick={() => navigate("/")}
        >
          <img src="/logo/mpbench.png" width={40} height={40} alt="Logo" />
          <h1 className="ml-4">MPBench</h1>
        </div>
      </div>
    </div>
  );
};

export default Header;
