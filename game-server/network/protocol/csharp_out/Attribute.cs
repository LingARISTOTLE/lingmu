// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: proto/attribute.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace MessageId {

  /// <summary>Holder for reflection information generated from proto/attribute.proto</summary>
  public static partial class AttributeReflection {

    #region Descriptor
    /// <summary>File descriptor for proto/attribute.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static AttributeReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChVwcm90by9hdHRyaWJ1dGUucHJvdG8SCW1lc3NhZ2VJZCofCg5BdHRyaWJ1",
            "dGVEYWlseRINCglOb25lRGFpbHkQACohCg9BdHRyaWJ1dGVXZWVrbHkSDgoK",
            "Tm9uZVdlZWtseRAAKiMKEEF0dHJpYnV0ZU1vbnRobHkSDwoLTm9uZU1vbnRo",
            "bHkQACodCg1BdHRyaWJ1dGVPbmNlEgwKCE5vbmVPbmNlEABCDFoKL2F0dHJp",
            "YnV0ZWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::MessageId.AttributeDaily), typeof(global::MessageId.AttributeWeekly), typeof(global::MessageId.AttributeMonthly), typeof(global::MessageId.AttributeOnce), }, null, null));
    }
    #endregion

  }
  #region Enums
  /// <summary>
  ///每日更新属性
  /// </summary>
  public enum AttributeDaily {
    [pbr::OriginalName("NoneDaily")] NoneDaily = 0,
  }

  /// <summary>
  ///每周更新属性
  /// </summary>
  public enum AttributeWeekly {
    [pbr::OriginalName("NoneWeekly")] NoneWeekly = 0,
  }

  /// <summary>
  ///每月更新属性
  /// </summary>
  public enum AttributeMonthly {
    [pbr::OriginalName("NoneMonthly")] NoneMonthly = 0,
  }

  /// <summary>
  ///一次更新属性
  /// </summary>
  public enum AttributeOnce {
    [pbr::OriginalName("NoneOnce")] NoneOnce = 0,
  }

  #endregion

}

#endregion Designer generated code
