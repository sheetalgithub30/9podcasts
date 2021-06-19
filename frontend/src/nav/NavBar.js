import { Layout, Avatar, Row } from "antd";
import { UserOutlined } from "@ant-design/icons";
import React from "react";

const NavBar = () => {
  return (
    <Layout.Header>
      <Row justify="space-between">
        <span style={{ color: "#ffffff", fontWeight: "bold", fontSize: 28 }}>
          9 Podcasts
        </span>
        <span style={{ display: "flex", alignItems: "center" }}>
          <span
            style={{
              color: "#ffffff",
              fontWeight: "bold",
              fontSize: 22,
              marginRight: 20,
            }}
          >
            My Podcasts
          </span>
          <Avatar size="large" icon={<UserOutlined />} />
        </span>
      </Row>
    </Layout.Header>
  );
};

export default NavBar;
