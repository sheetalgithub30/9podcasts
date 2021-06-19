import Layout from "antd/lib/layout/layout";
import React from "react";
import NavBar from "../nav/NavBar";
import CreatePodcast from "./CreatePodcast";
//import NoPodcast from "./NoPodcast";

const HomePage = () => {
  return (
    <Layout>
      <NavBar />
      <CreatePodcast />
    </Layout>
  );
};

export default HomePage;
