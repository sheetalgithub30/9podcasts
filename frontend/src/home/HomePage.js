import Layout from "antd/lib/layout/layout";
import React from "react";
import NavBar from "../nav/NavBar";
import NoPodcast from "./NoPodcast";

const HomePage = () => {
  return (
    <Layout>
      <NavBar />
      <NoPodcast />
    </Layout>
  );
};

export default HomePage;
