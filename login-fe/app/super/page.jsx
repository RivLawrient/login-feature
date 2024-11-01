// pages/index.js
"use client";
import { useRef, useEffect } from "react";

export default function Home() {
  const dragRef = useRef(null);
  const offset = useRef({ x: 0, y: 0 });

  const handleMouseDown = (e) => {
    offset.current = {
      x: e.clientX - dragRef.current.getBoundingClientRect().left,
      y: e.clientY - dragRef.current.getBoundingClientRect().top,
    };

    window.addEventListener("mousemove", handleMouseMove);
    window.addEventListener("mouseup", handleMouseUp);
  };

  const handleMouseMove = (e) => {
    dragRef.current.style.left = `${e.clientX - offset.current.x}px`;
    dragRef.current.style.top = `${e.clientY - offset.current.y}px`;
    dragRef.current.style.position = "absolute";
  };

  const handleMouseUp = () => {
    window.removeEventListener("mousemove", handleMouseMove);
    window.removeEventListener("mouseup", handleMouseUp);
  };

  useEffect(() => {
    // Cleanup event listeners on component unmount
    return () => {
      window.removeEventListener("mousemove", handleMouseMove);
      window.removeEventListener("mouseup", handleMouseUp);
    };
  }, []);

  return (
    <div style={{ position: "relative", height: "100vh" }}>
      <h1>Draggable Div Example</h1>
      <div
        ref={dragRef}
        onMouseDown={handleMouseDown}
        style={{
          width: "100px",
          height: "100px",
          backgroundColor: "lightblue",
          cursor: "grab",
          position: "absolute",
          top: "50px", // Starting position
          left: "50px", // Starting
        }}
      ></div>
    </div>
  );
}
