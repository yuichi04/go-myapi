package main_test

import (
	"fmt"
	"testing"
)

/*
テストの後処理が必要なケースには「データベース内の状態を現状復帰させる」パターンがある（テストの結果データベースに挿入されたレコードは、テスト終了時に削除する）

テストの後処理は t.Cleanupメソッドを使用する
t.Cleanupメソッドは引数が必要
引数には関数を渡す

t.Parallelメソッドを使用してサブテストを並列実行する際には、正しい順番で（=最後に）後処理実行するためにt.Cleanupメソッドで定義する必要がある。
*/

func TestA(t *testing.T) {
	fmt.Println("====================")

	// TestAにおける後処理の定義
	t.Cleanup(func() {
		// 後処理内容
		fmt.Println("cleanup")
	})

	// テスト実施
	fmt.Println("testA")
}

// 後処理に t.Cleanup メソッドを使わなければいけないケース
// t.Cleanup メソッドを使わずに後処理を実装しておかしくなる例
// 一番最初にdeferにセットした関数が実行される（"cleanup"が出力される）
func TestTableDrivienParallelA(t *testing.T) {
	fmt.Println("====================")

	// 後処理の記述
	defer func() {
		fmt.Println("cleanup")
	}()

	// 本来のテストの記述
	tests := []struct {
		testTitle string
	}{
		{testTitle: "subtest1"},
		{testTitle: "subtest2"},
		{testTitle: "subtest3"},
	}

	for _, test := range tests {
		testcase := test
		t.Run(testcase.testTitle, func(t *testing.T) {
			t.Parallel() // サブテストを並列で走らせる
			fmt.Println(testcase.testTitle)
		})
	}
}

// t.Cleanupメソッドを使用することで、最後に後処理が実行される例
func TestTableDrivienParallelB(t *testing.T) {
	fmt.Println("====================")

	// 後処理の記述
	t.Cleanup(func() {
		fmt.Println("cleanup")
	})

	// 本来のテストの記述
	tests := []struct {
		testTitle string
	}{
		{testTitle: "subtest1"},
		{testTitle: "subtest2"},
		{testTitle: "subtest3"},
	}

	for _, test := range tests {
		testcase := test
		t.Run(testcase.testTitle, func(t *testing.T) {
			t.Parallel() // サブテストを並列で走らせる
			fmt.Println(testcase.testTitle)
		})
	}
}
